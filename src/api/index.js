import axios from 'axios';

const BASE_URL = 'http://localhost:5000/api/users';

export const getPosts = async (page = 1, limit = 5, field = 'id', order = 'desc') => {
  try {
    const response = await axios.get(`${BASE_URL}?page=${page}&limit=${limit}&field=${field}&order=${order}`,
      {withCredentials: true,}
    );
    return response.data;
  } catch (error) {
    console.error('API Error:', error);
    throw error;
  }
};

export const getUserById = async (id) => {
  const res = await fetch(`http://localhost:5000/api/user/${id}`,{
    credentials: 'include',
  });
  return await res.json();
};

export const updateUser = async (id, userData) => {
  const res = await fetch(`http://localhost:5000/api/user/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(userData),
    credentials: 'include',
  });
  return await res.json();
};

export const loginUser = async (email, password) => {
  try{
    const response =  await axios.post('http://localhost:5000/api/login',{
      email,
      password,
    },{
      withCredentials: true,
    });
    return response.data;
  }catch(error){
    console.error('Login API Error: ', error.response?.data || error.message);
    throw error;
  }
};
