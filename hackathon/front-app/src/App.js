import React from 'react';
import './App.css';
import Form from "./Form.js"
import User from "./User.js"

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h2>User Register</h2>
      </header>
      <div className="App-main">
        <div className="App-form">
          <Form/>
        </div>
        <div className="App-user">
          <User/>
        </div>
      </div>
    </div>
  );

}

export default App;

