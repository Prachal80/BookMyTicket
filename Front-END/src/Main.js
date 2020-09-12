import React, { Component } from "react";
import { Route} from "react-router-dom";
import MyComponent from "./FirstPage";
import { Redirect } from 'react-router';
import PersonList from "./post";
import DisplayList from "./get";
import SendMail from "./mail"
import SendSms from "./sms"
import DeleteList from "./delete"
import UserApi from "./user"
import MovieApi from "./movie"
import TheaterApi from "./theater"
import BookApi from "./book"
import BookconfirmApi from "./bookconfirm"
import GetconfirmationApi from "./getconfirmation"
import home from "./home"
import Addmovie from "./addmovie";
import Addtheater from"./addtheater";
import Removetheater from "./removetheater";
import Removemovie from "./removemovie";
import Verifybooking from "./verify"

//Create a Main Component
class Main extends Component {
    render() {
        return (
            <div>
                {/*Render Different Component based on Route*/}

                <Route path="/" component={MyComponent} />
                {/* <Route path="/post" component={PersonList} />
                <Route path="/get" component={DisplayList} />
                <Route path="/mail" component={SendMail} />
               < Route path="/delete" component={DeleteList} />
               < Route path="/delete" component={DeleteList} />
                <Route path="/sms" component={SendSms} /> */}

                <Route path="/user" component={UserApi}/>
                <Route path="/movie" component={MovieApi}/>
                <Route path="/theater" component={TheaterApi}/>
                <Route path="/book" component={BookApi}/>
                <Route path="/bookconfirm" component={BookconfirmApi}/>
                <Route path="/getconfirmation" component={GetconfirmationApi}/>
                <Route path="/home" component={home}/>
                <Route path="/addtheater" component={Addtheater}/>
                <Route path="/addmovie" component={Addmovie}/>
                <Route path="/removetheater" component={Removetheater}/>
                <Route path="/removemovie" component={Removemovie}/>
                <Route path="/verify" component={Verifybooking}/>


            </div>
        );
    }
}
//Export The Main Component
export default Main;
