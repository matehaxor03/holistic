import TablePage from '../components/TablePage';
//import AppContext from '../AppContext';

export default class Dispatcher {
    pages = {"TablePage": TablePage};
    context = null;

    constructor(c) {
        this.context = c;
    }

    viewPage = (params) => {
        var Zlass = this.pages[params.type];
        var instance = <Zlass id={params.type} params={params}></Zlass>;
        this.context.updateState({...this.context, currentPage: instance});
      }
}