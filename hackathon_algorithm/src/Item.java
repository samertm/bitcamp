


public class Item {
    public String name;
    public int cal;
    public int fat;
    public int sugar;
    public int sodium;
    public int cost;

	public Item(String name, int cal, int fat, int sodium, int sugar, int cost) {
        this.name = name;
        this.cal = cal;
        this.fat = fat;
        this.sugar=sugar;
        this.sodium=sodium;
        this.cost=cost;
     }
        
        public int getCost() {
		return cost;
	}

	public void setCost(int cost) {
		this.cost = cost;
	}

		public String getName() {
    		return name;
    	}

    	public void setName(String name) {
    		this.name = name;
    	}

    	public int getCal() {
    		return cal;
    	}

    	public void setCal(int cal) {
    		this.cal = cal;
    	}

    	public int getFat() {
    		return fat;
    	}

    	public void setFat(int fat) {
    		this.fat = fat;
    	}

    	public int getSugar() {
    		return sugar;
    	}

    	public void setSugar(int sugar) {
    		this.sugar = sugar;
    	}

    	public int getSodium() {
    		return sodium;
    	}

    	public void setSodium(int sodium) {
    		this.sodium = sodium;
    	}
    
 
   
 
} // class
