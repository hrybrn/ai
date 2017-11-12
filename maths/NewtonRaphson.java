class NewtonRaphson {
    public static void main(String[] args){
        double current = 1;
        int i = 0;
        while(Math.abs(current) > 0.01){
            System.out.println(i++ + "\t" + current);

            current = current - (f(current) / fdash(current));
        }

        System.out.println(i++ + "\t" + current);
    }

    public static double f(double x){
        //x^3
        return Math.pow(x, 3);
    }

    public static double fdash(double x){
        //3x^2
        return 3 * Math.pow(x, 2); 
    }
}