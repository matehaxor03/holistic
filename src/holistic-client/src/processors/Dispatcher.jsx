import TablePage from '../components/TablePage';


export default class Dispatcher {
    pages = {"TablePage": TablePage};

    viewPage = (state, params) => {
        var Zlass = this.pages[params.type];
        var instance = <Zlass id={params.type} params={params}></Zlass>;
        
        state.updateState({...state, currentPage: instance});
        //console.log(context);
        //context.state.setState({...context.state.state, currentPage: instance});
        //context.ui.forceUpdate();
      }
}