/**
 * Created by ekeu on 05/04/16.
 */
import {bootstrap} from "angular2/platform/browser";
import {Component, Injectable} from "angular2/core";
import {Observable} from "rxjs";
import {Response, Http, HTTP_PROVIDERS} from "angular2/http";

// Service: People service retrieves the JSON data provided by the golang nameserver
@Injectable()
class PeopleSvc{
    constructor(private http: Http){} // Injects the Http service
    getPeople(): Observable<Response> {
        return this.http.get('http://localhost:8001/');
    }
}

@Component({
    selector : "app",
    templateUrl : 'tpls/app.html',
    providers : [HTTP_PROVIDERS, PeopleSvc]
}) class App{
    private name: string;
    private people: Observable<Response>;
    private num_persons : number;
    constructor(private peopleSvc : PeopleSvc){
        this.name = "Using HTTP in Angular2";
        peopleSvc.getPeople().subscribe(resp => this.people = resp.json());
        this.num_persons = 10;
    }

}

bootstrap(App, []);