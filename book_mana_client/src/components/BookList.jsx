import React, { useState, useEffect } from 'react';
import axios from 'axios';

const BookList = () => {
  const [books, setBooks] = useState([]);

  const fetchBooks = async () => {
    try{
        
        const response = await axios.get('http://localhost:8000/books');
        console.log(response.data);
        
        setBooks(response.data);
    }
    catch (error){
        alert('Error getting books');
        console.error(error);
    }
  };

  const handleDelete = async (id) => {
    try{
      console.log(id);
      
        await axios.delete(`http://localhost:8000/books/${id}`);
        fetchBooks();
    }
    catch (error){
        alert('Error deleting book');
        console.error(error);
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  return (
    <div>
      <h2>Book List</h2>
      <table>
        <thead>
          <tr>
            <th>Title</th>
            <th>Author</th>
            <th>Year</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {books.map((book) => (
            <tr key={book.id}>
              <td>{book.title}</td>
              <td>{book.author}</td>
              <td>{book.year}</td>
              <td>
                <button onClick={() => handleDelete(book.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default BookList;
