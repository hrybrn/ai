class NewtonRaphson {
    public static void main(String[] args){
        double current = 1;
        int i = 1;
        while(Math.abs(current) > 0.01){
            System.out.println(i++ + "\t" + current);

            current = current - (f(current) / fdash(current));
        }

        System.out.println(i++ + "\t" + current);
    }

    public static double f(double x){
        return Math.pow(x, 4);
    }

    public static double fdash(double x){
        //3x^2
        return 3 * Math.pow(x, 2); 
    }
}