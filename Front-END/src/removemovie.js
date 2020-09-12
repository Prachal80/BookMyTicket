import React from 'react';


export default class Removemovie extends React.Component {
  state = {
    name: '',
   
  }
  handleChange = event => {
    this.setState({ name: event.target.value });
  }
  
  

  handleSubmit = event => {
    try {
      event.preventDefault();
      const user = {
        name: this.state.name,
        
      };
      console.log(user);
      let result = fetch('http://34.68.43.210:8000/movieapi/movie/'+user.name,{
        method: 'post',
        mode: 'no-cors',
        headers:{
          'Accept': 'application/json',
          'Content-type': 'application/json'
        }
      });
      console.log("Result : "+result);
      window.location.replace("/movie");

    } catch (e) {
      console.log(e);
    }
    
  }

  render() {
    return (
      <div style={{padding: "40px"}} >
          <form onSubmit={this.handleSubmit} className="form-horizontal" style={{paddingLeft:"10%",paddingTop:"0%"}}>
          <label style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}} >Name of Movie you want to delete :</label><input className="form-control" type="text" id="name" onChange={this.handleChange} required="true" style={{width: "20%"}} ></input><br></br><br></br>
          
          <button style={{padding: "10px"}} className="btn btn-danger" type="submit">Click to Remove Movie</button><br></br><br></br>
          </form>
      </div>
    )
  }
}