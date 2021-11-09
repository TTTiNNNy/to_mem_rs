package main

import (
    "fmt"
    "strconv"
//     "bufio"
//     "io"
    "io/ioutil"
    "os"
)

var help_message = "to_mem.rs [OPTIONS] file ...\n\r" +
                           "-l \t\t set length for every part of data. e.g. for number 0x12345678 1: 0x12 0x34... 2: 0x1234 0x4567 etc. \n\r" +
                           "-o \t\t set name and path to output file. \n\r" +
                           "-f \t\t set input file format (current support only bin file)";
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func help_check(){
    //if(os.Args[2]){}
}

func is_need_help(){
    if(len(os.Args) > 2){
        if ((os.Args[1] != "-l") && (os.Args[1] != "-o") && (os.Args[1] != "-f")){
            fmt.Println(help_message)
            os.Exit(0)
        }
    }
}

func parse_argument(arg string, val string, mapped_args *map[string]string){
    switch
    {
        case arg == "-l":
        (*mapped_args)["-l"] = val

        case arg == "-o":
        (*mapped_args)["-o"] = val

        case arg == "-f":
        (*mapped_args)["-f"] = val

        case arg == "-h" || arg == "--help":
            fmt.Println(help_message)
            os.Exit(0)

        default:
            fmt.Println("to_mem_rs: " + arg + " is unknown command \n\r Run \"to_mem_rs --help\" or \"to_mem_rs -h\" for usage");
            os.Exit(0)
    }
}

func fill_mapped_args( mapped_args *map[string]string ){
        if(len(os.Args) > 2){
            for i := 2; i < len(os.Args); i+=2 {
                parse_argument(os.Args[i - 1], os.Args[i], mapped_args);
            }
        } else {
            if os.Args[1] == "-h" || os.Args[1] == "--help"{
                fmt.Println(help_message)
                os.Exit(0)
            } else {
                program_execute( mapped_args )
            }
        }
}

func verify_program( mapped_args *map[string]string ) (string, string, int) {

    var output_path string
    var word_length int
    var file_format string

    if o_flag, is_present := (*mapped_args)["-o"]; is_present != false{
        output_path = o_flag;
    } else {
        output_path = "mem_out.mem"
    }

    if l_flag, is_present := (*mapped_args)["-l"]; is_present != false{
        input_length, err := strconv.Atoi(l_flag)
        if  err != nil {
            fmt.Println("incorrect -l parameter. Terminate program")
            os.Exit(1)
        }
        word_length = input_length
    } else {
        word_length = 4;
    }

    if f_flag, is_present := (*mapped_args)["-f"]; is_present != false{
    if f_flag != "bin"{
        fmt.Println("incorrect -f parameter. Supported input formats: bin. Terminate program.")
        os.Exit(1)
    }
    file_format = f_flag;
    } else {
        file_format = "bin";
    }
    return output_path, file_format, word_length
}

func program_execute( mapped_args *map[string]string ){

    output_path, file_format, world_length := verify_program( mapped_args )

    var mem_string = "@00000000 "
    input_path, read_err := os.ReadFile(os.Args[len(os.Args) - 1]);
    if read_err != nil{
        fmt.Println("Cant create file. Terminate program")
        os.Exit(1)
    }
            for _, s := range input_path {
                fmt.Print(fmt.Sprintf("%x", s) + " ")
            }
            fmt.Println("");


    switch true{
        case file_format == "bin":
            for i, s := range input_path {
                if len(fmt.Sprintf("%x", s)) % 2 == 0{
                    mem_string += fmt.Sprintf("%x", s)
                } else{
                    mem_string += "0" + fmt.Sprintf("%x", s);
                }

                fmt.Println(i, world_length);
                if (i % world_length) == 0 && i != 0{
                mem_string += " "
                }
            }
    }

    file, err_create := os.Create(output_path);
    if err_create != nil{
        fmt.Println("Cant create file. Terminate program")
        os.Exit(1)
    }

    defer file.Close()

    err_write := ioutil.WriteFile(output_path, []byte(mem_string), 0644)

    if err_write != nil{
        fmt.Println("Cant write data to file. Terminate program")
        os.Exit(1)
        }

    fmt.Println("Success complete. output file data: \n\r" + mem_string);



}

func main() {
        fmt.Print("args: ")
        fmt.Println(len(os.Args))
        is_need_help();

        mapped_args := make(map[string]string)
        fill_mapped_args( &mapped_args )
        program_execute( &mapped_args )

}
