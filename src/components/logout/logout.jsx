import axios from "axios";


export const logoutUser = async () => {
  try{
    const response = await axios.get('http://localhost:5000/api/logout', {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    const errorMsg =
      error.response?.data?.message ||
      error.response?.data ||
      error.message ||
      'Logout failed';
    throw new Error(errorMsg);
  }
};
