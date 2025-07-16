import './Main.css';
import React, { useEffect, useState } from 'react';
import { getPosts } from '../../api';

const Main = () => {
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
    <main>
      {data.length > 0 ? (
        <table>
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
              <tr key={user._id || index}>
                <td data-label="S. No.">{index + 1}</td>
                <td data-label="Name">{user.name}</td>
                <td data-label="Email">{user.email}</td>
                <td data-label="Mobile">{user.mobile}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>No data</p>
      )}
    </main>
  );
};

export default Main;
