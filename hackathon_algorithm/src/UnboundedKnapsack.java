import java.text.*;
 
public class UnboundedKnapsack {
 
    protected Item []  items = {
        new Item("rice", 120,10,4,20,3),
        new Item("beans",90,12,2,7,4),
        new Item("carrots",100,5,3,4,2),
        new Item("celery",60,2,1,7,1)
                               
    		
        //new Item("rice", 10,10,4,20,3),
        //new Item("beans",12,12,2,7,4),
        //new Item("carrots",5,5,3,4,2),
        //new Item("celery",2,2,1,7,1)
    		
        //new Item("rice", 4,10,4,20,3),
        //new Item("beans",2,12,2,7,4),
        //new Item("carrots",3,5,3,4,2),
        //new Item("celery",1,2,1,7,1)
    						   
        //new Item("rice", 20,10,4,20,3),
        //new Item("beans",7,12,2,7,4),
        //new Item("carrots",4,5,3,4,2),
        //new Item("celery",7,2,1,7,1)
    };
    int    n = items.length; // the number of items
    Item constraints= new Item("price",2000,60,20,200,55); //max constraints for each catagory
    //Item constraints= new Item("price",60,60,20,200,50);
    //Item constraints= new Item("price",80,60,20,200,50);
    //Item constraints= new Item("price",200,60,20,200,50);
    Item bestcalitem= new Item("best",0,0,0,0,0);
    Item bestfatitem= new Item("best",0,0,0,0,0);
    Item bestsoditem= new Item("best",0,0,0,0,0);
    Item bestsugitem= new Item("best",0,0,0,0,0);
    
    
    int  []  maxItcal = new int [n];  // maximum number of items for cal constraint
    int  []  maxItfat = new int [n];
    int  []  maxItsod = new int [n];
    int  []  maxItsug = new int [n];
    
    
    int  []    ical = new int [n];  // current indexes of items for cal
    int  []    ifat = new int [n];
    int  []    isod = new int [n];
    int  []    isug = new int [n];
    
    
    int  [] bestcal = new int [n];  // best amounts for cal
    int  [] bestfat = new int [n];
    int  [] bestsod = new int [n];
    int  [] bestsug = new int [n];
    
 
    public UnboundedKnapsack() {
        // initializing:
        for (int i = 0; i < n; i++) {
            maxItcal [i] = Math.min(
                                    (int)(constraints.getCal() / items[i].getCal()),//2000 is the cal constraint
                                    (int)(constraints.getCost() / items[i].getCost())// 50 in the cost contraint
                                    );
            maxItfat [i] = Math.min(
                                    (int)(constraints.getFat() / items[i].getFat()),//2000 is the cal constraint
                                    (int)(constraints.getCost() / items[i].getCost())// 50 in the cost contraint
                                    );
            maxItsod [i] = Math.min(
                                    (int)(constraints.getSodium() / items[i].getSodium()),//2000 is the cal constraint
                                    (int)(constraints.getCost() / items[i].getCost())// 50 in the cost contraint
                                    );
            maxItsug [i] = Math.min(
                                    (int)(constraints.getSugar() / items[i].getSugar()),//2000 is the cal constraint
                                    (int)(constraints.getCost() / items[i].getCost())// 50 in the cost contraint
                                    );
           
        } // for (i)
 
        // calc the solution:
        calcWithRecursion(0);
 
        // Print out the solution:
        NumberFormat nf = NumberFormat.getInstance();
        System.out.println("Maximum value achievable for cals is: " + bestcalitem.getCal());
        System.out.print("This is achieved by carrying (one solution): ");
        for (int i = 0; i < n; i++) {
            System.out.print(bestcal[i] + " " + items[i].getName() + ", ");
            System.out.print(bestfat[i] + " " + items[i].getName() + ", ");
            System.out.print(bestsod[i] + " " + items[i].getName() + ", ");
            System.out.print(bestsug[i] + " " + items[i].getName() + ", ");
            
        }
        System.out.println();
        System.out.println("The calories you will have is: " + nf.format(bestcalitem.getCal()) +
                           "   and the total amount used is: " + nf.format(bestcalitem.getCost())
                           );
        System.out.println("The calories you will have is: " + nf.format(bestfatitem.getCal()) +
                           "   and the total amount used is: " + nf.format(bestfatitem.getCost())
                           );
        System.out.println("The calories you will have is: " + nf.format(bestsoditem.getCal()) +
                           "   and the total amount used is: " + nf.format(bestsoditem.getCost())
                           );
        System.out.println("The calories you will have is: " + nf.format(bestsugitem.getCal()) +
                           "   and the total amount used is: " + nf.format(bestsugitem.getCost())
                           );
 
    }
 
    // calculation the solution with recursion method
    // item : the number of item in the "items" array
    public void calcWithRecursion(int item) {
        for (int i = 0; i <= maxItcal[item]; i++) {
            ical[item] = i;
            if (item < n-1) {
                calcWithRecursion(item+1);
            } else {
                int    currVal = 0;   // current value
                int currWei = 0; // current weight
                int currVol = 0; // current Volume
                int totalcost=0;
                for (int j = 0; j < n; j++) {
                    currVal += ical[j] * items[j].getCal();
                    currWei += ical[j] * items[j].getCal();
                    currVol += ical[j] * items[j].getCost();
                }
 
                if (currVal > constraints.getCost()
                    &&
                    currVol <= constraints.getCost()
                    )
                    {
                        bestcalitem.setCal (currVal);
                        bestcalitem.setCal(currWei);
                        bestcalitem.setCost(currVol);
                        totalcost= bestcalitem.cost;
                        for (int j = 0; j < n; j++) bestcal[j] = ical[j];
                    } // if (...)
            } // else
        } // for (i)
        for (int i = 0; i <= maxItfat[item]; i++) {
            ifat[item] = i;
            if (item < n-1) {
                calcWithRecursion(item+1);
            } else {
                int    currVal = 0;   // current value
                int currWei = 0; // current weight
                int currVol = 0; // current Volume
                
                for (int j = 0; j < n; j++) {
                    currVal += ifat[j] * items[j].getFat();
                    currWei += ifat[j] * items[j].getFat();
                    currVol += ifat[j] * items[j].getCost();
                }
 
                if (currVal > constraints.getCost()
                    &&
                    currVol <= constraints.getCost()
                    )
                    {
                        bestfatitem.setFat (currVal);
                        bestfatitem.setFat(currWei);
                        bestfatitem.setCost(currVol);
                    
                        for (int j = 0; j < n; j++) bestfat[j] = ifat[j];
                    } // if (...)
            } // else
        } // for (i)
        for (int i = 0; i <= maxItsod[item]; i++) {
            isod[item] = i;
            if (item < n-1) {
                calcWithRecursion(item+1);
            } else {
                int    currVal = 0;   // current value
                int currWei = 0; // current weight
                int currVol = 0; // current Volume
                int totalcost=0;
                for (int j = 0; j < n; j++) {
                    currVal += isod[j] * items[j].getSodium();
                    currWei += isod[j] * items[j].getSodium();
                    currVol += isod[j] * items[j].getCost();
                }
 
                if (currVal > constraints.getCost()
                    &&
                    currVol <= constraints.getCost()
                    )
                    {
                        bestsoditem.setSodium (currVal);
                        bestsoditem.setSodium(currWei);
                        bestsoditem.setCost(currVol);
                        totalcost= bestsoditem.cost;
                        for (int j = 0; j < n; j++) bestsod[j] = isod[j];
                    } // if (...)
            } // else
        } // for (i)
        for (int i = 0; i <= maxItsug[item]; i++) {
            isug[item] = i;
            if (item < n-1) {
                calcWithRecursion(item+1);
            } else {
                int    currVal = 0;   // current value
                int currWei = 0; // current weight
                int currVol = 0; // current Volume
                int totalcost=0;
                for (int j = 0; j < n; j++) {
                    currVal += isug[j] * items[j].getSugar();
                    currWei += isug[j] * items[j].getSugar();
                    currVol += isug[j] * items[j].getCost();
                }
 
                if (currVal > constraints.getCost()
                    &&
                    currVol <= constraints.getCost()
                    )
                    {
                        bestsugitem.setSugar (currVal);
                        bestsugitem.setSugar(currWei);
                        bestsugitem.setCost(currVol);
                        totalcost= bestsugitem.cost;
                        for (int j = 0; j < n; j++) bestsug[j] = isug[j];
                    } // if (...)
            } // else
        } // for (i)
    } // calcWithRecursion()
    // the main() function:
    public static void main(String[] args) {
        new UnboundedKnapsack();
    } // main()
} // class
