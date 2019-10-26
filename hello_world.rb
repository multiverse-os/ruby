puts "hello world"

if 3 > 5
  p "wrong"
else
  p "right"
end 

def Module
  def Class
    attr_accessor :hello

    def hello_world
      p "hi world"
    end 
  end 
end 


Module::Class.hello_world
