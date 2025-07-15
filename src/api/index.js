export const getPosts = async () => {
  const response = await fetch('http://localhost:5000/api/users');
  const data = await response.json();
  console.log("Fetched posts:", data);
  return data;
};
