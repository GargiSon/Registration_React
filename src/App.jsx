import { useEffect, useState } from 'react';
import './App.css';
import { getPosts } from './api';
import PostCard from './components/PostCard';

function App() {
  const [data, setData] = useState([]);

  useEffect(() => {
    getPosts()
      .then((posts) => {
        console.log("Posts from API:", posts);
        setData(posts);
      })
      .catch((err) => console.error("Fetch failed:", err));
  }, []);

  return (
    <div className="App">
      <h2>User List</h2>
      {data.length > 0 ? (
        <table border="1" cellPadding="10" cellSpacing="0">
          <thead>
            <tr>
              <th>S. No.</th>
              <th>Name</th>
              <th>Email</th>
              <th>Mobile</th>
            </tr>
          </thead>
          <tbody>
            {data.map((user, index) => (
              <PostCard
                key={user._id || index}
                index={index}
                name={user.name}
                email={user.email}
                mobile={user.mobile}
              />
            ))}
          </tbody>
        </table>
      ) : (
        <p>No data</p>
      )}
    </div>
  );
}

export default App;
