import React from 'react';
import { Route, Redirect } from 'react-router';
import MyComponent from "./FirstPage";
import Main from './Main';
import { BrowserRouter } from 'react-router-dom';
import swal from 'sweetalert';


export default class PersonList extends React.Component {

    constructor(props){
        super(props);
        this.state= {
           
            userloaded: false,
            movieloaded: false,
            theaterloaded: false

        }

    }


  state = {
    email: '',
    movie: '',
    theater: '',
    
  }
  componentDidMount(){
    fetch('http://34.68.43.210:8000/userapi/user/a')
    .then(res => res.json())
    .then(json => {
       console.log(json);
    });
}
  handleChange = event => {
    this.setState({ email: event.target.value });
  }
  handleChange2 = event => {
    this.setState({ movie: event.target.value });
  }
  handleChange3 = event => {
    this.setState({ theater: event.target.value });
  }
  handleSubmit = event => {
    const u = {
        user: this.state.email,
        movie: this.state.movie,
        theater: this.state.theater,
        
      };
      console.log(u);

        fetch('http://34.68.43.210:8000/userapi/user/'+u.user).then(function(response) {
        console.log(response.status); // returns 200
        if(response.status == 200){
            console.log("OK user found.....");


            fetch('http://34.68.43.210:8000/movieapi/movie/'+u.movie).then(function(response) {
            console.log(response.status); // returns 200
                if(response.status == 200){
                    console.log("OK movie found.....");


                        fetch('http://34.68.43.210:8000/theaterapi/theater/'+u.theater).then(function(response) {
                            console.log(response.status); // returns 200
                            if(response.status == 200){
                                console.log("OK theater found.....");
                                
                                //Redirect
                                // return( <Redirect to="/"></Redirect>);
                                window.location.replace("/bookconfirm/?user="+u.user+"&movie="+u.movie+"&theater="+u.theater);

                            }
                            else{
                                console.log("ERROR Match theater not found.....");
                                swal("ERROR!", "No such theater!", "error");
                            }

                        });
                    }
                else{
                    console.log("ERROR Match movie not found.....");
                    swal("ERROR!", "No such movie!", "error");
                }
            });
            
        }
        else{
            console.log("ERROR Match user not found.....");
            swal("ERROR!", "No such user!", "error");
        }
        // console.log("User Loaded  :  "+this.state.userloaded)
        });
  }
  render() {
    return (
        <div style={{padding: "40px"}}>
            <div  className="form-horizontal" style={{paddingLeft:"10%",paddingTop:"0%"}}>
            
        <label style={{padding: "10px"}} style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}}>Email    :</label><input type="email" className="form-control" style={{width: "20%"}} required="true" id="name" onChange={this.handleChange}></input><br></br><br></br>
        <label style={{padding: "10px"}} style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}}>Movie    :</label><input type="text" className="form-control" style={{width: "20%"}} required="true" id="email" onChange={this.handleChange2} ></input><br></br><br></br>
        <label style={{padding: "10px"}} style={{padding: "10px"}} style={{color:"#212529",fontSize:"xx-large"}}>Theater  :</label><input type="text"  className="form-control" style={{width: "20%"}} required="true" id="mobile" onChange={this.handleChange3}></input><br></br><br></br>
        <button onClick={this.handleSubmit} className="btn btn-success" type="submit" style={{padding: "10px"}}>Click to Book</button><br></br><br></br>
        </div>
        </div>
    )
  }
}








