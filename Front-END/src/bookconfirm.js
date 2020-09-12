import React from 'react';


export default class BookconfirmApi extends React.Component {

    constructor(props){
        super(props);
        this.state= {
            user: JSON,
            movie: JSON,
            theater: JSON,
            userloaded: false,
            movieloaded: false,
            theaterloaded: false,
        }
    }
    componentDidMount(){
        const query = new URLSearchParams(this.props.location.search);
          
            fetch('http://34.68.43.210:8000/userapi/user/'+query.get('user'))
            .then(res => res.json())
            .then(json => {
                this.setState({
                    userloaded: true,
                    user: json,
                })
            });

            fetch('http://34.68.43.210:8000/movieapi/movie/'+query.get('movie'))
            .then(res => res.json())
            .then(json => {
                this.setState({
                    movieloaded: true,
                    movie: json,
                })
            });

            fetch('http://34.68.43.210:8000/theaterapi/theater/'+query.get('theater'))
            .then(res => res.json())
            .then(json => {
                this.setState({
                    theaterloaded: true,
                    theater: json,
                })
            });
        
        
    }

    handleSubmit = event => {

        const user2 = {
            Name: this.state.user.name,
            email: this.state.user.email,
            movie: this.state.movie.name,
            theater: this.state.theater.name
          };
    
        let result = fetch('http://54.237.2.113:3000/person',{
        method: 'post',
        mode: 'no-cors',
        headers:{
          'Accept': 'application/json',
          'Content-type': 'application/json'
        },
        body: JSON.stringify(user2)
      });
      console.log("Result : "+result);
        
     window.location.replace("/verify/?mail="+this.state.user.email+"&movie="+this.state.movie.name+"&theater="+this.state.theater.name);
        
      }
  render() {
    var { userloaded, user, movie, theater } = this.state;

    if(!userloaded){
        return <div style={{padding: "40px",paddingTop:"0px",paddingLeft:"33%",width:"120%"}}>Loading User Details ...</div>
    }
    else{
        return (
        

            <div style={{padding: "40px",paddingTop:"70px",paddingLeft:"10%",width:"-webkit-fill-available"}}>
                
                <h3 style={{paddingLeft:"2.4%"}}> <span className="badge badge-success">User Details</span></h3>
                <br></br>
                 <ul>
                       
                            <li key={user._id} class="list-group-item list-group-item-action "style={{padding: "5px",width:"100%",fontSize: "xx-large",backgroundColor:"#F8F9FA"}}>
                            <span style={{color:"#0C5DBB",paddingRight:"50px"}} >   Name:</span>  {user.name}<br></br>
                            <span style={{color:"#0C5DBB",paddingRight:"50px"}} >    Email:</span> {user.email}<br></br>
                            </li>
                      
                 </ul>
                 <br></br>
                 <br></br>

                 <h3 style={{paddingLeft:"2.4%"}}> <span className="badge badge-success">Movie Details</span></h3>
                 <br></br>
                 <ul>
                       
                 <li key={movie._id} class="list-group-item list-group-item-action "style={{padding: "5px",width:"100%",fontSize: "xx-large",backgroundColor:"#F8F9FA"}}>
                 <span style={{color:"#0C5DBB",paddingRight:"90px"}} >   Name: </span>{movie.name}<br></br>
                 <span style={{color:"#0C5DBB",paddingRight:"68px"}} >    Ratings:</span> {movie.rating}<br></br>
                 <span style={{color:"#0C5DBB",paddingRight:"64px"}}>  Director:</span>  {movie.director}<br></br>
                             <span style={{color:"#0C5DBB",paddingRight:"100px"}}>  Stars:</span>  {movie.stars}<br></br>
                             <span style={{color:"#0C5DBB",paddingRight:"15px"}}>  Description:</span>  {movie.desc}<br></br>
                            </li>
                      
                 </ul>
                <br></br>
                <br></br>
                 <h3 style={{paddingLeft:"2.4%"}}> <span className="badge badge-success">Theater Details</span></h3>
                 <br></br>
                 <ul>
                       
                 <li key={theater._id} class="list-group-item list-group-item-action "style={{padding: "5px",width:"100%",fontSize: "xx-large",backgroundColor:"#F8F9FA"}}>
                 <span style={{color:"#0C5DBB",paddingRight:"50px"}} >   Name:</span> {theater.name}<br></br>
                 <span style={{color:"#0C5DBB",paddingRight:"20px"}} >   Address: </span>{theater.address}<br></br>
                 <span style={{color:"#0C5DBB",paddingRight:"27px"}} >    Ratings: </span>{theater.rating}<br></br>
                 <span style={{color:"#0C5DBB",paddingRight:"33px"}} >    Screen: </span>{theater.screens}<br></br>
                            </li>
                      
                 </ul>
                 <br></br>
                 <br></br>

                 <button style={{padding: "10px",marginLeft:"2.5%",width:"17%"}} className="btn btn-success" onClick={this.handleSubmit} type="submit">Click to Confirm Booking</button>
             
            </div>
          )
    }
    
   
  }
}