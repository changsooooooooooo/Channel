package InOut;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.Executor;
import java.util.concurrent.Executors;
import java.util.concurrent.ThreadFactory;
import java.util.stream.Collectors;

class FileReadTest {

    private final Executor executor =
            Executors.newCachedThreadPool(
                    new ThreadFactory() {
                        @Override
                        public Thread newThread(Runnable r) {
                            Thread t = new Thread(r);
                            t.setDaemon(true);
                            return t;
                        }
                    });

    @Test
    @DisplayName("BufferedReader Check")
    void brReadLineCheck() throws IOException {
        FileRead fr = FileRead.makeFileRead("../Source-Text/hamlet.txt", "hello");
        long start = System.currentTimeMillis();
        long answer = fr.countIncludeLines("hamlet");
        long end = System.currentTimeMillis();
        System.out.printf("processing time : %s%n", end - start);
        System.out.println(answer);
    }

    @Test
    @DisplayName("BufferedReader CompletableFuture Check")
    void brReadLineCheckCompletableFuture() throws IOException {
        FileRead fr = FileRead.makeFileRead("../Source-Text/hamlet.txt", "hello");
        List<String> brLines = fr.brLines;
        long start = System.currentTimeMillis();

        List<CompletableFuture<Boolean>> futureList = brLines.stream()
                .map(line -> CompletableFuture.supplyAsync(()->line.contains("hamlet"), executor))
                .collect(Collectors.toList());

        long answer = futureList.stream()
                .map(CompletableFuture::join)
                .filter(line -> line)
                .count();

        long end = System.currentTimeMillis();
        System.out.printf("processing time executor : %s%n", end - start);
        System.out.println(answer);
    }

}
