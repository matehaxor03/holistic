import React from 'react';
import { useDispatch } from 'react-redux'
import { viewPage } from '../reducers/AppSlice'

export default function ViewTableButton (props) { 
   const dispatch = useDispatch();

   return (
      <button type="button" className="btn btn-light" onClick={() => dispatch(viewPage({"type": "TablePage", "tableName": props.text}))}>{props.text}</button>
      );
   
}
