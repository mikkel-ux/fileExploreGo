export namespace goFiles {
	
	export class FileData {
	    name: string;
	    path: string;
	    size: string;
	    extension: string;
	    created: string;
	    modified: string;
	    accessed: string;
	    type: string;
	    isHidden: boolean;
	    isReadOnly: boolean;
	    base64?: string;
	    firstFrame?: string;
	
	    static createFrom(source: any = {}) {
	        return new FileData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.size = source["size"];
	        this.extension = source["extension"];
	        this.created = source["created"];
	        this.modified = source["modified"];
	        this.accessed = source["accessed"];
	        this.type = source["type"];
	        this.isHidden = source["isHidden"];
	        this.isReadOnly = source["isReadOnly"];
	        this.base64 = source["base64"];
	        this.firstFrame = source["firstFrame"];
	    }
	}
	export class ImageResponse {
	    data: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new ImageResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.type = source["type"];
	    }
	}
	export class dirData {
	    name: string;
	    path: string;
	    points: number;
	
	    static createFrom(source: any = {}) {
	        return new dirData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.points = source["points"];
	    }
	}

}

export namespace main {
	
	export class FileData {
	    name: string;
	    path: string;
	    size: number;
	    extension: string;
	    created: string;
	    modified: string;
	    accessed: string;
	    fileType: string;
	    permissions: number;
	    isHidden: boolean;
	    isReadOnly: boolean;
	    base64?: string;
	
	    static createFrom(source: any = {}) {
	        return new FileData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.size = source["size"];
	        this.extension = source["extension"];
	        this.created = source["created"];
	        this.modified = source["modified"];
	        this.accessed = source["accessed"];
	        this.fileType = source["fileType"];
	        this.permissions = source["permissions"];
	        this.isHidden = source["isHidden"];
	        this.isReadOnly = source["isReadOnly"];
	        this.base64 = source["base64"];
	    }
	}

}

