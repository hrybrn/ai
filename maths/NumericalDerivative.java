class NumericalDerivative {
    public static void main(String[] args){

        double x = 0.5;

        for(int i = 0; i < 20; i++){
            double h = Math.pow(10, -i);

            System.out.println(h + "\t" + g(x, h));
        }
    }

    public static double f(double x){
        return Math.exp((Math.pow(Math.cos(x), 2)) + Math.tan(x) * Math.exp(-(Math.pow(x,2))));
    }

    public static double g(double x, double h){
        return h / (f(x + h) / f(x));
    }
}