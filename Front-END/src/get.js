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
        fetch('http://ubunto-go-clb-1702965941.us-east-1.elb.amazonaws.com/people')
        .then(res => res.json())
        .then(json => {
            this.setState({
                isLoaded: true,
                items: json,
            })
        });
    }
  render() {
    var { isLoaded, items } = this.state;

    if(!isLoaded){
        return <div>Loading ...</div>
    }
    else{
        return (
        

            <div>
              
              <ul>
                        {items.map(item => (
                            <li key={item._id} style={{paddingBottom: "20px"}}>
                                User Id: {item._id}<br></br>
                                Name: {item.name}<br></br>
                                Email: {item.email}<br></br>
                                Mobile: {item.mobile}<br></br>
                            </li>
                        ))}
                    </ul>
             
            </div>
          )
    }
   
  }
}