import React from 'react';


export default class MovieApi extends React.Component {

    constructor(props){
        super(props);
        this.state= {
            items: [],
            isLoaded: false
        }
    }
    componentDidMount(){
        fetch('http://34.68.43.210:8000/movieapi/movies')
        .then(res => res.json())
        .then(json => {
            this.setState({
                isLoaded: true,
                items: json,
            })
        });
    }

    handleSubmit1 = event => {
        window.location.replace("/addmovie");
    }

    handleSubmit2 = event => {
        window.location.replace("/removemovie");
    }

  render() {
    var { isLoaded, items } = this.state;

    if(!isLoaded){
        return <div>Loading ...</div>
    }
    else{
        return (
        

            <div  style={{padding: "40px",paddingTop:"0px",paddingLeft:"10%",width:"-webkit-fill-available"}}>
              <button onClick={this.handleSubmit1} style={{padding: "10px",marginLeft:"40px"}} className="btn btn-success" type="submit">Add Movie</button>
                    <button onClick={this.handleSubmit2} style={{padding: "10px",marginLeft:"40px"}} className="btn btn-danger" type="submit">Remove Movie</button><br></br><br></br>
              <ul>
                        {items.map(item => (
                            <li  class="list-group-item list-group-item-action "style={{padding: "5px",width:"100%",fontSize: "xx-large",backgroundColor:"#F8F9FA"}}>
                               
                             <span style={{color:"#0C5DBB",paddingRight:"90px"}}> Name:</span>   {item.name}<br></br>
                             <span style={{color:"#0C5DBB",paddingRight:"68px"}}>  Ratings:</span>  {item.rating}<br></br>
                             <span style={{color:"#0C5DBB",paddingRight:"64px"}}>  Director:</span>  {item.director}<br></br>
                             <span style={{color:"#0C5DBB",paddingRight:"100px"}}>  Stars:</span>  {item.stars}<br></br>
                             <span style={{color:"#0C5DBB",paddingRight:"15px"}}>  Description:</span>  {item.desc}<br></br>
                                
                            </li>
                        ))}
                    </ul>
                    
            </div>
          )
    }
   
  }
}