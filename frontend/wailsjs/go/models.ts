export namespace factorial {
	
	export class EvolutaESL {
	    Mean: number;
	    Max: number;
	    Min: number;
	    Median: number;
	    Variance: number;
	    StandardDeviation: number;
	    Skewness: number;
	    Values: number[];
	
	    static createFrom(source: any = {}) {
	        return new EvolutaESL(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mean = source["Mean"];
	        this.Max = source["Max"];
	        this.Min = source["Min"];
	        this.Median = source["Median"];
	        this.Variance = source["Variance"];
	        this.StandardDeviation = source["StandardDeviation"];
	        this.Skewness = source["Skewness"];
	        this.Values = source["Values"];
	    }
	}

}

