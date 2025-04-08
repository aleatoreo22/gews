export namespace model {
	
	export class News {
	    Link: string;
	    Headline: string;
	    SubHeadling: string;
	    ImageLink: string;
	
	    static createFrom(source: any = {}) {
	        return new News(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Link = source["Link"];
	        this.Headline = source["Headline"];
	        this.SubHeadling = source["SubHeadling"];
	        this.ImageLink = source["ImageLink"];
	    }
	}
	export class NewsSite {
	    Site: string;
	    News: News[];
	    Visible: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NewsSite(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Site = source["Site"];
	        this.News = this.convertValues(source["News"], News);
	        this.Visible = source["Visible"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

