import React from "react";
import { useState } from "react";
// import fetchUsers  from "./User";

const Form = () => {
  const [name, setName] = useState("");
  const [age, setAge] = useState(0);
  const onSubmit = async (event) => {
    event.preventDefault();
    if (!name) {
      alert("Please enter name");
      return;
    }

    if (name.length > 50) {
      alert("Please enter a name shorter than 50 characters");
      return;
    }

    if (age < 20 || age > 80) {
      alert("Please enter age between 20 and 80");
      return;
    }

    try {
      const result = await fetch("http://localhost:8000/user", {
        method: "POST",
        // headers: {
        //   "Content-Type": "application/json"
        // },
        body: JSON.stringify({
          name: name,
          age: age,
        }),
      });
      if (!result.ok) {
        throw Error(`Failed to create user: ${result.status}`);
      }

      setName("");
      setAge(0);
      // fetchUsers();
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <form style={{ display: "flex", flexDirection: "column" }}>
      <label>Name: </label>
      <input
        type={"text"}
        value={name}
        onChange={(e) => setName(e.target.value)}
      ></input>
      <label>Age: </label>
      <input
        type={"number"}
        value={age}
        onChange={(e) => setAge(e.target.valueAsNumber)}
      ></input>
      <button onClick={onSubmit}>POST</button>
    </form>
  );
};

export default Form;