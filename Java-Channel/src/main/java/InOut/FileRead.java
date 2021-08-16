package InOut;

import java.io.*;
import java.util.ArrayList;
import java.util.List;

public class FileRead {

    public final List<String> brLines;

    private FileRead(List<String> brLines){
        this.brLines = brLines;
    }

    public static FileRead makeFileRead(String path, String fileName) throws IOException {
        File file = new File(path+fileName);
        FileInputStream fis = new FileInputStream(file);
        BufferedReader br = new BufferedReader(new InputStreamReader(fis));
        List<String> brLines = new ArrayList<>();
        brLines = makeStringArray(br);
        return new FileRead(brLines);
    }

    private static List<String> makeStringArray(BufferedReader br) throws IOException {
        List<String > brLines = new ArrayList<>();
        String string = br.readLine();
        while(string != null){
            brLines.add(string);
            string = br.readLine();
        }
        return brLines;
    }

    public long countIncludeLines(String word) throws IOException {
        return brLines.stream()
                .filter(line->line.contains(word))
                .count();
    }
}
