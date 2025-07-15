const PostCard = ({ index, name, email, mobile }) => {
  return (
    <tr>
      <td>{index + 1}</td>
      <td>{name}</td>
      <td>{email}</td>
      <td>{mobile}</td>
    </tr>
  );
};

export default PostCard;
