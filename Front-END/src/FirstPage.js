import React, { Component } from "react";

import { Link } from "react-router-dom";


export default class MyComponent extends Component {
    

    render() {
       
            return (
                <div  style={{paddingLeft:"15%",paddingBottom:"5%"}}>
                    
                    <Link to={`/home`} style={{padding: "40px"}} > <button class="btn btn-primary" type="submit" style={{padding: "20px",width:"10%"}}>HOME</button></Link>
                    <Link to={`/user`} style={{padding: "40px"}} > <button class="btn btn-primary" type="submit" style={{padding: "20px",width:"10%"}}>User</button></Link>
                    <Link to={`/movie`} style={{padding: "40px"}} > <button class="btn btn-primary" type="submit" style={{padding: "20px",width:"10%"}}>Movie</button></Link>
                    <Link to={`/theater`}  style={{padding: "40px"}}> <button class="btn btn-primary" type="submit" style={{padding: "20px",width:"10%"}}>Theater</button></Link>
                    <Link to={`/book`} style={{padding: "40px"}}> <button class="btn btn-primary" type="submit" style={{padding: "20px",width:"10%"}}>Book</button></Link>
                  
                   
                </div>
                
            );
    }
}