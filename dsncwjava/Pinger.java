import java.io.InputStream;
import java.util.Scanner;

public class Pinger {
    public static String headings = "Host, IPv4, Avg RTT, Min RTT, Max RTT, IPv6, Avg RTT, Min RTT, Max RTT";

    public static String[] websites = {
        "google.com",
        "google.co.uk",
        "google.com.au",
        "youtube.com",
        "instagram.com",
        "yandex.ru",
        "wikipedia.org",
        "stackoverflow.com",
        "linkedin.com",
        "yahoo.com",
        "yahoo.co.uk",
        "yahoo.com.au",
        "facebook.com",
        "netflix.com",
        "ebay.co.uk",
        "bbc.co.uk",
        "www.bbc.co.uk",
        "reddit.com",
        "github.com",
        "taobao.com",
        "mail.ru",
        "wikia.com",
        "loopsofzen.uk",
        "iinet.net.au"};
    
    public static void main(String[] args){
        Runtime rt = Runtime.getRuntime();

        System.out.println(headings);

        for(int i = 0; i < websites.length; i++){
            try{
                Process pr = rt.exec("ping4 -c 10 " + websites[i]);
                InputStream err = pr.getErrorStream();
                InputStream is = pr.getInputStream();
                String output = convertStreamToString(is);
                
                String errors = convertStreamToString(err).trim();
                if(errors.equals("")){
                
                    String[] lines = output.split("\n");
                    String[] stats = lines[lines.length - 1].split("=")[1].split("/");
                    
                    String ip = lines[2].split("\\(")[1].split("\\)")[0];


                    for(int j = 0; j < stats.length; j++){
                        stats[j] = stats[j].trim().replace("ms", "");
                    }

                    //print the ipv4
                    System.out.print(websites[i] + "," + ip + "," + stats[1] + "," + stats[0] + "," + stats[2] + ",");
                } else {
                    //print errors in place if it fails
                    System.out.print(websites[i] + ",err,err,err,err,");
                }

                Process pr6 = rt.exec("ping6 -c 10 " + websites[i]);
                InputStream err6 = pr6.getErrorStream();
                InputStream is6 = pr6.getInputStream();
                String output6 = convertStreamToString(is6);
                
                String errors6 = convertStreamToString(err6).trim();
                if(errors6.equals("")){
                
                    String[] lines = output6.split("\n");
                    String[] stats = lines[lines.length - 1].split("=")[1].split("/");
                    String ip = lines[1].split("\\(")[1].split("\\)")[0];
                    
                    for(int j = 0; j < stats.length; j++){
                        stats[j] = stats[j].trim().replace("ms", "");
                    }

                    //print the ipv6
                    System.out.println(ip + "," + stats[1] + "," + stats[0] + "," + stats[2] + ",");
                } else {
                    //print errors in place if it fails
                    System.out.println("err,err,err,err");
                }
            } catch(Exception e){
                e.printStackTrace();
            }
        }
    }

    static String convertStreamToString(InputStream is) {
        Scanner s = new Scanner(is);
        s.useDelimiter("\\A");

        String out = "";
        if(s.hasNext()){
            out = s.next();
        }

        s.close();
        return out;
    }
}