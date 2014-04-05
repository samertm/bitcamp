






import java.text.*;
 
public class holder {
 
    protected Item []  items = {
                               new Item("rice", 120,10,4,20,3),
                               new Item("beans",90,12,2,7,4),
                               new Item("carrots",100,5,3,4,2),
                               new Item("celery",60,2,1,7,1)
                               };
    int    n = items.length; // the number of items
    Item constraints= new Item("price",2000,60,20,200,50); //max constraints for each catagory
    Item bestcalitem= new Item("best",0,0,0,0,0);
    Item bestfatitem= new Item("best",0,0,0,0,0);
    Item bestsugitem= new Item("best",0,0,0,0,0);
    Item bestcarbitem= new Item("best",0,0,0,0,0);
    int  []  maxItcal = new int [n];  // maximum number of items for cal constraint
    int  []  maxItfat = new int [n];
    int  []  maxItsug = new int [n];
    int  []  maxItcarb = new int [n];
    int  []    ical = new int [n];  // current indexes of items for cal
    int  []    ifat = new int [n];
    int  []    isug = new int [n];
    int  []    icarb = new int [n];
    int  [] bestcal = new int [n];  // best amounts for cal
    int  [] bestfat = new int [n];
    int  [] bestsug = new int [n];
    int  [] bestcarb = new int [n];
 
    public holder() {
        // initializing:
        for (int i = 0; i < n; i++) {
            maxItcal [i] = Math.min(
                           (int)(2000 / items[i].getCal()),//2000 is the cal constraint
                           (int)(50 / items[i].getCost())// 50 in the cost contraint
                        );
            maxItfat [i] = Math.min(
                    (int)(60 / items[i].getFat()),
                    (int)(50 / items[i].getCost())// 50 in the cost contraint
                 );
            maxItsug [i] = Math.min(
                    (int)(20 / items[i].getSugar()),//2000 is the cal constraint
                    (int)(50 / items[i].getCost())// 50 in the cost contraint
                 );
            maxItcarb [i] = Math.min(
                    (int)(200 / items[i].getCarbs()),//2000 is the cal constraint
                    (int)(50 / items[i].getCost())// 50 in the cost contraint
                 );
        } // for (i)
 
        // calc the solution:
        calcWithRecursion(0);
 
        // Print out the solution:
        NumberFormat nf = NumberFormat.getInstance();
        System.out.println("Maximum value achievable for cals is: " + bestcalitem.getCal());
        System.out.println("Maximum value achievable for fat is: " + bestfatitem.getFat());
        System.out.println("Maximum value achievable for sugar is: " + bestsugitem.getSugar());
        System.out.println("Maximum value achievable for carbs is: " + bestcarbitem.getCarbs());
        System.out.print("This is achieved by carrying (one solution): ");
        for (int i = 0; i < n; i++) {
            System.out.print(bestcal[i] + " " + items[i].getName() + ", ");
            System.out.print(bestfat[i] + " " + items[i].getName() + ", ");
            System.out.print(bestsug[i] + " " + items[i].getName() + ", ");
            System.out.print(bestcarb[i] + " " + items[i].getName() + ", ");
        }
        System.out.println();
        System.out.println("The calories you will have is: " + nf.format(bestcalitem.getCal()) +
                           "   and the total amount used is: " + nf.format(bestcalitem.getCost())
                          );
        System.out.println("The fat you will have is: " + nf.format(bestfatitem.getCal()) +
                "   and the total amount used is: " + nf.format(bestfatitem.getCost())
               );
        System.out.println("The sugar you will have is: " + nf.format(bestsugitem.getCal()) +
                "   and the total amount used is: " + nf.format(bestsugitem.getCost())
               );
        System.out.println("The carbs you will have is: " + nf.format(bestcarbitem.getCal()) +
                "   and the total amount used is: " + nf.format(bestcarbitem.getCost())
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
                for (int j = 0; j < n; j++) {
                    currVal += ical[j] * items[j].getCal();// bug here!!!!!
                    currWei += ical[j] * items[j].getCal();//bug here!!!!!
                    currVol += ical[j] * items[j].getCost();
                }
 
                if (currVal > constraints.getCost()
                    &&
                    currWei <= constraints.getCal()
                    &&
                    currVol <= 2000
                )
                {
                    bestcalitem.setCal (currVal);// bug here!!!!!
                    bestcalitem.setCal(currWei);// bug here!!!!!
                    bestcalitem.setCost(currVol);
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
                    currVal += ifat[j] * items[j].getFat();// bug here!!!!!
                    currWei += ifat[j] * items[j].getFat();//bug here!!!!!
                    currVol += ifat[j] * items[j].getCost();
                }
 
                if (currVal > constraints.getCost()
                    &&
                    currWei <= constraints.getFat()
                    &&
                    currVol <= 60
                )
                {
                    bestfatitem.setFat (currVal);// bug here!!!!!
                    bestfatitem.setFat(currWei);// bug here!!!!!
                    bestfatitem.setCost(currVol);
                    for (int j = 0; j < n; j++) bestfat[j] = ifat[j];
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
                for (int j = 0; j < n; j++) {
                    currVal += isug[j] * items[j].getSugar();// bug here!!!!!
                    currWei += isug[j] * items[j].getSugar();//bug here!!!!!
                    currVol += isug[j] * items[j].getSugar();
                }
 
                if (currVal > constraints.getCost()
                    &&
                    currWei <= constraints.getSugar()
                    &&
                    currVol <= 20
                )
                {
                    bestsugitem.setSugar (currVal);// bug here!!!!!
                    bestsugitem.setSugar(currWei);// bug here!!!!!
                    bestsugitem.setCost(currVol);
                    for (int j = 0; j < n; j++) bestsug[j] = isug[j];
                } // if (...)
            } // else
        } // for (i)
        
        for (int i = 0; i <= maxItcarb[item]; i++) {
            icarb[item] = i;
            if (item < n-1) {
                calcWithRecursion(item+1);
            } else {
                int    currVal = 0;   // current value
                int currWei = 0; // current weight
                int currVol = 0; // current Volume
                for (int j = 0; j < n; j++) {
                    currVal += icarb[j] * items[j].getCarbs();// bug here!!!!!
                    currWei += icarb[j] * items[j].getCarbs();//bug here!!!!!
                    currVol += icarb[j] * items[j].getCost();
                }
 
                if (currVal > constraints.getCost()
                    &&
                    currWei <= constraints.getCarbs()
                    &&
                    currVol <= 200
                )
                {
                    bestcarbitem.setCarbs (currVal);// bug here!!!!!
                    bestcarbitem.setCarbs(currWei);// bug here!!!!!
                    bestcarbitem.setCost(currVol);
                    for (int j = 0; j < n; j++) bestcarb[j] = icarb[j];
                } // if (...)
            } // else
        } // for (i)
        
        
    } // calcWithRecursion()
 
    // the main() function:
    public static void main(String[] args) {
        new holder();
    } // main()
 
} // class
