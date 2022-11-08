import React from "react";
import {useEffect, useState} from "react";
const User= () => {
    const [name, setName] = useState("");
    const [age, setAge] = useState(0);
    const [id, setId] = useState();
    const [users, setUsers] = useState([]);
    useEffect(() => {
      const fetchUsers = async () => {
        try {
          const res = await fetch("http://localhost:8000/user",{
            method:"GET",
            headers: {
               "Content-Type": "application/json",
             },
            params: {
                id:id,
                name:name,
                age: age
           },
        });
          if (!res.ok) {
            throw Error(`Failed to fetch users: ${res.status}`);
          }
          const user = await res.json();
          console.log(user)
          setUsers(user);
        } catch (err) {
          console.error(err);
        }
    };
    fetchUsers();
  },[]);

      return (
        <div className="Users">
            <ul className="User-main">
                   {users.map((User)=>{
                       return <li key={User.id}>{User.name}{User.age}</li>
                   })}
            </ul>
        </div>
      );
};

export default User;

