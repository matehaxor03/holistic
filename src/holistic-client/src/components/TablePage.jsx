import ContentPage from './ContentPage'

export default class TablePage extends ContentPage { 
   render() {
      return <h1>{JSON.stringify(this.props.params)}</h1>;
   }
}