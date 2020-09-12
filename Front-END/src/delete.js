import React from 'react';


export default class DeleteList extends React.Component {
  state = {
    id: '',
  }
  handleChange = event => {
    this.setState({ id: event.target.value });
  }
  

  handleSubmit = event => {
    try {
        event.preventDefault();
        
        let result = fetch('http://34.68.43.210:8000/theaterapi/theater/Shiv',{
          method: 'delete',
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
          <label>Id of Element you want to Delete  </label><input type="text" id="id" onChange={this.handleChange}></input>
          <button onClick={this.handleSubmit} type="submit">Click to post data</button>
       
      </div>
    )
  }
}