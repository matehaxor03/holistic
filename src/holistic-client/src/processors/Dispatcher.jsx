import TablePage from '../components/TablePage';


export default class Dispatcher {
    pages = {"TablePage": TablePage};

    viewPage = (context, params) => {
        var Zlass = this.pages[params.type];
        var instance = <Zlass id={params.type} params={params}></Zlass>;
        console.log(context);
        context.ui.setState({...context.ui.state, currentPage: instance});
        //context.ui.forceUpdate();
      }
}