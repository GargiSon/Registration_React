import React from 'react';

export const logoutUser = async () => {
  const response = await fetch('http://localhost:5000/api/logout', {
    method: 'GET', 
    credentials: 'include',
  });

  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || 'Logout failed');
  }

  return response.json();
};
