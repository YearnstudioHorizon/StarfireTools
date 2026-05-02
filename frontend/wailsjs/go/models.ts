export namespace main {
	
	export class ItemGroup {
	    name: string;
	    items: string[];
	    items_descs: string[];
	
	    static createFrom(source: any = {}) {
	        return new ItemGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.items = source["items"];
	        this.items_descs = source["items_descs"];
	    }
	}

}

