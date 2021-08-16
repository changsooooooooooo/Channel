package InOut;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class Files {

    private final List<FileRead> frList;

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

}
