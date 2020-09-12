import React from 'react';


export default class DisplayList extends React.Component {

    constructor(props){
        super(props);
        this.state= {
            items: [],
            isLoaded: false
        }
    }
    componentDidMount(){
        fetch('http://34.68.43.210:8000/theaterapi/theaters')
        .then(res => res.json())
        .then(json => {
            this.setState({
                isLoaded: true,
                items: json,
            })
        });
    }


    handleSubmit1 = event => {
        window.location.replace("/addtheater");
    }

    handleSubmit2 = event => {
        window.location.replace("/removetheater");
    }

  render() {
    var { isLoaded, items } = this.state;

    if(!isLoaded){
        return <div>Loading ...</div>
    }
    else{
        return (
        

            <div style={{padding: "40px",paddingTop:"0px",paddingLeft:"10%",width:"-webkit-fill-available"}}>
              <button onClick={this.handleSubmit1} style={{padding: "10px",marginLeft:"40px"}} className="btn btn-success" type="submit">Add Theater</button>
                    <button onClick={this.handleSubmit2}  style={{padding: "10px",marginLeft:"40px"}} className="btn btn-danger" type="submit">Remove Theater</button><br></br><br></br>
              <ul>
                        {items.map(item => (
                            <li class="list-group-item list-group-item-action " key={item._id} style={{padding: "5px",width:"40%",fontSize: "xx-large",backgroundColor:"#F8F9FA"}}>
                               <span style={{color:"#0C5DBB",paddingRight:"50px"}} > Name:</span> {item.name}<br></br>
                               <span style={{color:"#0C5DBB",paddingRight:"20px"}} >Address:</span>  {item.address}<br></br>
                               <span style={{color:"#0C5DBB",paddingRight:"27px"}} > Ratings: </span> {item.rating}<br></br>
                               <span style={{color:"#0C5DBB",paddingRight:"33px"}} > Screen: </span> {item.screens}<br></br>
                            </li>
                        ))}
                    </ul>
                    
            </div>
          )
    }
   
  }
}