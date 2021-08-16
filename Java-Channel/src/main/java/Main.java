import InOut.Files;

import java.io.IOException;

public class Main {
    public static void main(String[] args) throws IOException {
        Files files = Files.makeFiles("../Source-Text/", "*");
        Long answer = files.showResult("romeo");
        System.out.println("Result : " + answer);
    }
}
