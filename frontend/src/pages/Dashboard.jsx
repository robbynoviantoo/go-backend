import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../services/api';

const Dashboard = () => {
  const [items, setItems] = useState([]);
  const [users, setUsers] = useState([]);
  const [me, setMe] = useState(null);
  const [error, setError] = useState('');
  const [newItemName, setNewItemName] = useState('');
  const [newItemRemarks, setNewItemRemarks] = useState('');
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  const fetchData = async () => {
    try {
      const [itemsRes, usersRes, meRes] = await Promise.all([
        api.get('/api/items'),
        api.get('/api/users'),
        api.get('/api/me')
      ]);
      setItems(itemsRes.data.data || []);
      setUsers(usersRes.data || []);
      setMe(meRes.data || null);
    } catch (err) {
      if (err.response?.status === 401) {
        localStorage.removeItem('token');
        navigate('/login');
      }
      console.error('Error fetching data:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [navigate]);

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  const handleAddItem = async (e) => {
    e.preventDefault();
    if (!newItemName.trim()) return;
    
    try {
      await api.post('/api/items', { 
        name: newItemName,
        remarks: newItemRemarks 
      });
      setNewItemName('');
      setNewItemRemarks('');
      fetchData(); // Refresh list
    } catch (err) {
      console.error('Error adding item:', err);
    }
  };

const handleDeleteItem = async (id) => {
  if (!window.confirm('Are you sure you want to delete this item?')) return;

  try {
    await api.delete(`/api/items/${id}`);
    fetchData();
  } catch (err) {
    if (err.response) {
      setError(err.response.data.error || 'Failed to delete item');
    } else {
      setError('Network error');
    }
  }
};

  if (loading) {
    return <div className="auth-container"><div className="auth-title">Loading...</div></div>;
  }

  return (
    <div className="dashboard-container">
      <div className="dashboard-header">
        <h2 className="auth-title" style={{ margin: 0, fontSize: '1.75rem' }}>Dashboard</h2>
        <h2 className="auth-title" style={{ margin: 0, fontSize: '1.75rem' }}>Welcome, {me?.name || 'Loading...'}</h2>
        <button onClick={handleLogout} className="btn btn-secondary">Log Out</button>
      </div>

      <div className="card">
        <h3 className="card-title">Manage Items</h3>
        
        <form onSubmit={handleAddItem} className="add-item-form">
          <div className="form-group">
            <input 
              className="form-input" 
              type="text" 
              value={newItemName}
              onChange={(e) => setNewItemName(e.target.value)}
              placeholder="Item Name"
              required
            />
          </div>
          <div className="form-group">
            <input 
              className="form-input" 
              type="text" 
              value={newItemRemarks}
              onChange={(e) => setNewItemRemarks(e.target.value)}
              placeholder="Remarks (optional)"
            />
          </div>
          <button type="submit" className="btn">Add Item</button>
        </form>
        {error && (
          <div style={{ color: 'red', marginBottom: '10px' }}>
            {error}
          </div>
        )}
        <div className="item-list">
          {items.map(item => (
            <div key={item.id} className="item-card">
              <div className="item-info">
                <h4>{item.name}</h4>
                {item.remarks && <p>{item.remarks}</p>}
                <p style={{ fontSize: '0.75rem', marginTop: '4px', opacity: 0.7 }}>
                  Added by User ID: {item.user_id}
                </p>
              </div>
              <button 
                onClick={() => handleDeleteItem(item.id)}
                className="btn btn-danger"
              >
                Delete
              </button>
            </div>
          ))}
          {items.length === 0 && (
            <div style={{ textAlign: 'center', color: 'var(--text-muted)' }}>
              No items available.
            </div>
          )}
        </div>
      </div>

      <div className="card">
        <h3 className="card-title">Registered Users</h3>
        <div className="item-list">
          {users.map(user => (
            <div key={user.id} className="item-card">
              <div className="item-info">
                <h4>{user.name}</h4>
                <p>{user.email}</p>
              </div>
            </div>
          ))}
        </div>
      </div>

    </div>
  );
};

export default Dashboard;
