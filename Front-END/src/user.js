import React from 'react';


export default class UserApi extends React.Component {
  state = {
    name: '',
    email: '',
    password: '',
  }
  handleChange = event => {
    this.setState({ name: event.target.value });
  }
  handleChange2 = event => {
    this.setState({ email: event.target.value });
  }
  handleChange3 = event => {
    this.setState({ password: event.target.value });
  }
  

  handleSubmit = event => {
    try {
      event.preventDefault();
      const user = {
        name: this.state.name,
        email: this.state.email,
        password: this.state.password
      };
      console.log(user);
      let result = fetch('http://34.68.43.210:8000/userapi/user',{
        method: 'post',
        mode: 'no-cors',
        headers:{
          'Accept': 'application/json',
          'Content-type': 'application/json'
        },
        body: JSON.stringify(user)
      });
      console.log("Result : "+result);
      

    } catch (e) {
      console.log(e);
    }
    
  }

  render() {
    return (
      <div style={{padding: "40px"}} >
          <form onSubmit={this.handleSubmit} className="form-horizontal" style={{paddingLeft:"10%",paddingTop:"0%"}}>
          <label style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}} >Name     :</label><input className="form-control" type="text" id="name" onChange={this.handleChange} required="true" style={{width: "20%"}} ></input><br></br><br></br>
          <label style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}} >Email :</label><input className="form-control" type="email" id="email" onChange={this.handleChange2} required="true" style={{width: "20%"}}></input><br></br><br></br>
          <label style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}} >Password :</label><input className="form-control" type="password" id="password" onChange={this.handleChange3} required="true" style={{width: "20%"}}></input><br></br><br></br>
          <button style={{padding: "10px"}} className="btn btn-success" type="submit">Click to Create User</button><br></br><br></br>
          </form>
      </div>
    )
  }
}