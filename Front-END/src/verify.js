import React from 'react';


export default class Verifybooking extends React.Component {

    constructor(props){
        super(props);
        this.state= {
            items: [],
            isLoaded: false,
            users: [],
        }
    }
    componentDidMount(){
        const query = new URLSearchParams(this.props.location.search);
        fetch('http://34.68.43.210:8000/showapi/show/'+query.get('theater')+query.get('movie')).then(function(response) {
        console.log(response.status); // returns 200
        console.log(response);
        if(response.status != 200){
           
           

              const user3 = {
                ShowID: query.get('theater')+query.get('movie'),
                TheatreID: query.get('theater'),
                MovieID: query.get('movie')

               
              };
              console.log("--------------Creating show-----------");
              console.log(user3)

              let result3 = fetch('http://34.68.43.210:8000/showapi/show',{
                method: 'post',
                   mode: 'no-cors',
               headers:{
                   'Accept': 'application/json',
                   'Content-type': 'application/json'
               },
           body: JSON.stringify(user3)
         });
         console.log("------------Show Created----------------");
         console.log("Result : "+result3);

         console.log("------------Booking Entry----------------");
         const user = {
            ShowID: query.get('theater')+query.get('movie'),
           
          };
          console.log(user);

            let result = fetch('http://34.68.43.210:8000/showapi/createbook/'+query.get('theater')+query.get('movie'),{
             method: 'post',
                mode: 'no-cors',
            headers:{
                'Accept': 'application/json',
                'Content-type': 'application/json'
            },
        body: JSON.stringify(user)
      });
      console.log("------------Bokk Entry Done----------------")
      console.log("Result : "+result);
      

     
            
        }
    //     else{
    //         //console.log("ERROR Match user not found.....");
    //     }
    //     // console.log("User Loaded  :  "+this.state.userloaded)
    const user2 = {
        ShowID: query.get('theater')+query.get('movie'),
        User: query.get('mail')
       
      };
      console.log("--------------Entering User details-----------");
      console.log(user2);
    let result2 = fetch('http://34.68.43.210:8000/showapi/book/'+query.get('theater')+query.get('movie'),{
         method: 'post',
            mode: 'no-cors',
        headers:{
            'Accept': 'application/json',
            'Content-type': 'application/json'
        },
    body: JSON.stringify(user2)
  });
  console.log("--------------User Entered -----------");
  console.log("Result : "+result2);

     });

       




     window.location.replace("/getconfirmation/?mail="+query.get('mail')+"&movie="+query.get('movie')+"&theater="+query.get('theater'));
        
    }


    handleSubmit1 = event => {
        window.location.replace("/addtheater");
    }

    handleSubmit2 = event => {
        window.location.replace("/removetheater");
    }

  render() {
      let users = this.state.items.map(item => {
        return(
            <div>
            <span style={{color:"#0C5DBB",paddingRight:"20px"}} >{item.Users} </span><br></br>
            </div>
        )
      })

    var { isLoaded, items } = this.state;

    if(!isLoaded){
        return <div>Loading ...</div>
    }
    else{
        return (
        

            <div style={{padding: "40px",paddingTop:"0px",paddingLeft:"10%",width:"-webkit-fill-available"}}>
               <ul>
                        {items.map(item => (
                            <li class="list-group-item list-group-item-action " key={item._id} style={{padding: "5px",width:"40%",fontSize: "xx-large",backgroundColor:"#F8F9FA"}}>
                               <span style={{color:"#0C5DBB",paddingRight:"50px"}} > Name:</span> {item.ShowID}<br></br>
                               {users}
                               {/* <span style={{color:"#0C5DBB",paddingRight:"20px"}} >Address:</span>  {item.Users}<br></br> */}
                              
                            </li>
                        ))}
                    </ul>
                    
            </div>
          )
    }
   
  }
}