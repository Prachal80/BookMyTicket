import React from 'react';


export default class PersonList extends React.Component {
  state = {
    name: '',
    email: '',
    mobile: '',
  }
  handleChange = event => {
    this.setState({ name: event.target.value });
  }
  handleChange2 = event => {
    this.setState({ email: event.target.value });
  }
  handleChange3 = event => {
    this.setState({ mobile: event.target.value });
  }
  

  handleSubmit = event => {
    try {
      event.preventDefault();
      const user = {
        UserId: this.state.UserId,
        name: this.state.name,
        email: this.state.email,
        mobile: this.state.mobile
      };
      console.log(user);
      let result = fetch('http://ubunto-go-clb-1702965941.us-east-1.elb.amazonaws.com/person',{
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
      <div>
          <label>Name</label><input type="text" id="name" onChange={this.handleChange}></input>
          <label>Email</label><input type="text" id="email" onChange={this.handleChange2}></input>
          <label>Mobile Number (add with code)</label><input type="text" id="mobile" onChange={this.handleChange3}></input>
          <button onClick={this.handleSubmit} type="submit">Click to post data</button>
       
      </div>
    )
  }
}