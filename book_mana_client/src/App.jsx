import React, { useState } from 'react';
import BookList from './components/BookList';
import BookForm from './components/BookForm';

const App = () => {
  const [currentBook, setCurrentBook] = useState(null);

  const handleEdit = (book) => {
    setCurrentBook(book);
  };

  const handleSave = () => {
    setCurrentBook(null);
  };

  return (
    <div>
      <h1>Book Management</h1>
      <BookForm currentBook={currentBook} onSave={handleSave} />
      <BookList onEdit={handleEdit} />
    </div>
  );
};

export default App;
