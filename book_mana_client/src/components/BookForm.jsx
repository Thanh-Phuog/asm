import React, { useState, useEffect } from 'react';
import axios from 'axios';
const BookForm = ({ currentBook, onSave }) => {
  const [book, setBook] = useState({ title: '', author: '', year: '' });

  useEffect(() => {
    if (currentBook) {
      setBook(currentBook);
    }
  }, [currentBook]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setBook({ ...book, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
        if (book.id) {
            await axios.put(`http://localhost:8000/books/${book.id}`, book);
        } else {
            await axios.post('http://localhost:8000/books', book);
        }
        onSave();
        setBook({ title: '', author: '', year: '' });
        

    }
    catch (error) {
        alert('Error Saving Book');
        console.error(error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>{book.id ? 'Edit Book' : 'Add New Book'}</h2>
      <input
        type="text"
        name="title"
        placeholder="Title"
        value={book.title}
        onChange={handleChange}
        required
      />
      <input
        type="text"
        name="author"
        placeholder="Author"
        value={book.author}
        onChange={handleChange}
        required
      />
      <input
        type="number"
        name="year"
        placeholder="Year"
        value={book.year}
        onChange={handleChange}
      />
      <button type="submit">{book.id ? 'Update' : 'Create'}</button>
    </form>
  );
};

export default BookForm;
