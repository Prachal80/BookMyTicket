import React from 'react';


export default class Addtheater extends React.Component {
  state = {
    name: '',
    address: '',
    rating: '',
    screen: '',
  }
  handleChange = event => {
    this.setState({ name: event.target.value });
  }
  handleChange2 = event => {
    this.setState({ address: event.target.value });
  }
  handleChange3 = event => {
    this.setState({ rating: event.target.value });
  }
  handleChange4 = event => {
    this.setState({ screen: event.target.value });
  }
  

  handleSubmit = event => {
    try {
      event.preventDefault();
      const user = {
        name: this.state.name,
        address: this.state.address,
        rating: this.state.rating,
        screens: this.state.screen
      };
      console.log(user);
      let result = fetch('http://34.68.43.210:8000/theaterapi/theater',{
        method: 'post',
        mode: 'no-cors',
        headers:{
          'Accept': 'application/json',
          'Content-type': 'application/json'
        },
        body: JSON.stringify(user)
      });
      console.log("Result : "+result);
      window.location.replace("/theater");

    } catch (e) {
      console.log(e);
    }
    
  }

  render() {
    return (
      <div style={{padding: "40px"}} >
           <h3 style={{paddingLeft:"10%",paddingBottom:"1%"}}> <span className="badge badge-success">Enter Theater Details</span></h3>
          <form onSubmit={this.handleSubmit} className="form-horizontal" style={{paddingLeft:"10%",paddingTop:"0%"}}>
          <label style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}} >Name     :</label><input className="form-control" type="text" id="name" onChange={this.handleChange} required="true" style={{width: "20%"}} ></input><br></br><br></br>
          <label style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}} >Address :</label><input className="form-control" type="text" id="email" onChange={this.handleChange2} required="true" style={{width: "20%"}}></input><br></br><br></br>
          <label style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}} >Rating :</label><input className="form-control" type="text" id="password" onChange={this.handleChange3} required="true" style={{width: "20%"}}></input><br></br><br></br>
          <label style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}} >Screens :</label><input className="form-control" type="text" id="password" onChange={this.handleChange4} required="true" style={{width: "20%"}}></input><br></br><br></br>
          <button style={{padding: "10px"}} className="btn btn-success" type="submit">Click to Add Theater</button><br></br><br></br>
          </form>
      </div>
    )
  }
}