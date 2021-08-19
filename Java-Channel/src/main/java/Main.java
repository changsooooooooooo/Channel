import InOut.Files;

import java.io.IOException;

public class Main {
    public static void main(String[] args) throws IOException {
        Files files = Files.makeFiles("./Source-Text/", "*");
        long start = System.currentTimeMillis();
        Long answer = files.showResult("this");
        long end = System.currentTimeMillis();
        System.out.printf("Result : %d Elapsed Time : %d", answer, end-start);
    }
}
