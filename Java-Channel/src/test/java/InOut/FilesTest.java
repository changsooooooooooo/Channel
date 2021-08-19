package InOut;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.List;

class FilesTest {

    @Test
    @DisplayName("Performance Comparing")
    void showResultTest() throws IOException {
        Files files = Files.makeFiles("../Source-Text/", "*");
        long start = System.currentTimeMillis();
        Long answer = files.showResult("this");
        long end = System.currentTimeMillis();
        System.out.println("answer : " + answer);
        System.out.printf("Processing Time Completable Future : %s", end-start);
        System.out.println();
    }

    @Test
    @DisplayName("Performance Comparing Just Sequential")
    void showResultTest2() throws IOException {
//        Files files = Files.makeFiles("../Source-Text/", "*");
//        List<FileRead> frList = files.frList;
//        long start = System.currentTimeMillis();
//        Long answer = frList.stream()
//                .mapToLong(file -> file.countIncludeLines("this"))
//                .sum();
//        long end = System.currentTimeMillis();
//        System.out.println("answer : " + answer);
//        System.out.printf("Processing Time : %s", end-start);
    }
}
