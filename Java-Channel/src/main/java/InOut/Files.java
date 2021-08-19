package InOut;

import executor.ExecutorVariable;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.Executor;
import java.util.stream.Collectors;

public class Files {

    public final List<FileRead> frList;

    private final Executor executorVariable = new ExecutorVariable().excutable;

    private Files(List<FileRead> frList){
        this.frList = frList;
    }

    public static Files makeFiles(String path, String fileName) throws IOException {
        List<String> fileNameList = new ArrayList<>();
        fileNameList = makeFileNameList(path, fileName);

        List<FileRead> frList = new ArrayList<>();
        for(String file : fileNameList){
            FileRead fr = FileRead.makeFileRead(path, file);
            frList.add(fr);
        }
        return new Files(frList);
    }

    private static List<String> makeFileNameList(String path, String fileName){

        if(fileName.equals("*")){
            File file = new File(path);
            return List.of(Objects.requireNonNull(file.list()));
        }
        return List.of(fileName.split(" "));
    }

    public Long showResult(String word) {
        List<CompletableFuture<Long>> futureList = frList.stream()
                .map(file -> CompletableFuture.supplyAsync(()-> file.countIncludeLines(word), executorVariable))
                .collect(Collectors.toList());

        return futureList.stream()
                .mapToLong(CompletableFuture::join)
                .sum();
    }
}
