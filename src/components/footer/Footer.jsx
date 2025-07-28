import React from "react";
import './Footer.css';

export default function Footer(props){
    return(
        <div className="footer">
            <h1>{props.note}</h1>
        </div>
    );
}