import React from 'react';


export default class SendMail extends React.Component {
  state = {
    id: '',
   
  }
  handleChange = event => {
    this.setState({ id: event.target.value });
  }
  
  

  handleSubmit = event => {
    try {
      event.preventDefault();
      
      let result = fetch('http://localhost:12345/sendmail/'+ this.state.id,{
        method: 'post',
        mode: 'no-cors',
        headers:{
          'Accept': 'application/json',
          'Content-type': 'application/json'
        }
      });
      console.log("Result : "+result);

    } catch (e) {
      console.log(e);
    }
    
  }

  render() {
    return (
      <div>
          <label>Enter your account ID </label><input type="text" id="id" onChange={this.handleChange}></input>
          
          <button onClick={this.handleSubmit} type="submit">Click to send mail</button>
       
      </div>
    )
  }
}