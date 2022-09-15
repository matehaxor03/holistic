import TablePage from '../components/TablePage';


export default class Dispatcher {
    pages = {"TablePage": TablePage};

    viewPage = (context, pageName, params) => {
        var Zlass = this.pages[pageName];
        var instance = <Zlass id={pageName} params={params}></Zlass>;
        console.log(context);
        context.state = {...context.state, currentPage: instance};
        context.ui.forceUpdate();
      }
}