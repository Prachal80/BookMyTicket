import React from 'react';
import swal from 'sweetalert';
import { watchFile } from 'fs';
import { delay } from 'q';


export default class GetconfirmationApi extends React.Component {
  state = {
    email: '',
    mobile: '',
  }
  handleChange2 = event => {
    this.setState({ email: event.target.value });
  }
  handleChange3 = event => {
    this.setState({ mobile: event.target.value });
  }
  
  componentDidMount(){
    const query = new URLSearchParams(this.props.location.search);
    this.setState({ email: query.get('mail') });
  }

  handleSubmit = event => {
    try {
      event.preventDefault();
      const user = {
        email: this.state.email,
        mobile: this.state.mobile
      };
      console.log(user);
      console.log(this.state.email)
      console.log(this.state.mobile)
     // swal("GREAT!", "Enjoy the movie!", "success");
      let result = fetch('http://54.237.2.113:3000/sendmail/'+this.state.email,{
        method: 'post',
        mode: 'no-cors',
        headers:{
          'Accept': 'application/json',
          'Content-type': 'application/json'
        }
      });
      console.log("Result : "+result);


   

      let result2 = fetch('http://54.237.2.113:3000/sendsms/'+this.state.email+','+this.state.mobile,{
        method: 'post',
        mode: 'no-cors',
        headers:{
          'Accept': 'application/json',
          'Content-type': 'application/json'
        }
      });
      console.log("Result : "+result2);

    swal("GREAT!", "Enjoy the movie!", "success")
    .then((value) => {
      if(value != null){
       window.location.replace("/home");
      }
    });


    } catch (e) {
      console.log(e);
    }
    
  }

  render() {
    return (
        <div style={{padding: "40px"}}>
            <h1  style={{padding: "40px",paddingTop:"0px",paddingLeft:"10%",width:"-webkit-fill-available"}}> <span className="badge badge-success">Get your E-ticket...</span></h1>
        <div  className="form-horizontal" style={{paddingLeft:"10%"}}>
        <h3  style={{padding: "40px",paddingTop:"0px",paddingLeft:"0%",width:"-webkit-fill-available"}}> <span className="badge badge-warning">Verify your E-mail address ...</span></h3>
    <label style={{padding: "10px"}} style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}}>Email    :</label><input value={this.state.email} type="email" className="form-control" style={{width: "20%"}} required="true" id="name" onChange={this.handleChange2}></input><br></br><br></br>
    <h3  style={{padding: "40px",paddingTop:"0px",paddingLeft:"0%",width:"-webkit-fill-available"}}> <span className="badge badge-warning">Also want a SMS ?   Enter your Number... </span></h3>
    <label style={{padding: "10px"}} style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}}>Mobile Number    :</label><input type="text" className="form-control" style={{width: "20%"}} required="true" id="email" onChange={this.handleChange3} ></input><br></br><br></br>
    <button onClick={this.handleSubmit} className="btn btn-success" type="submit" style={{padding: "10px"}}>Get E-tickets</button><br></br><br></br>
    </div>
    </div>
    )
  }
}