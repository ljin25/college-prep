/**
 * A program that reads a file containing information about Amazon's
 * top 50 books from years 2009 to 2019, then takes the ratings of
 * each book, displays a histogram, and calculates the mean, median,
 * and standard deviation of the data.
 *
 * @author Layton Jin
 * @version 20211214
 *
 */
 
import java.util.Scanner;
import java.util.ArrayList;
import java.util.List;
import java.util.Collections;
import java.io.File;
import java.io.FileNotFoundException;

public class Analysis {
   
    /**
     * Reads a file, initializes each line into a String array using
     * split() around each comma of the line, then adds the contents
     * of what is at the index containing the user rating in each line.
     * If there is an error reading the file, prints "ERROR - File fName
     * not found" where fName is the name of the file.
     *
     * @param fName
     *            the String name of the file to be read
     * @return a Double List representing the ratings read from the file
     */
   public static List<Double> getRatings(String fName) {
      
      //declare returning Double List
      ArrayList<Double> ratings = new ArrayList<Double>();
      try {
         
         //create File and Scanner objects
         File file = new File(fName);
         Scanner read = new Scanner(file);
         
         //initialize to first line of file to skip the "sample" line
         String fileLine = read.nextLine();
         String[] column;
         
         while (read.hasNext()) {
            fileLine = read.nextLine();
            
            //create String array for each line of file and split by commas
            column = fileLine.split(",");
            
            //convert contents of third column of String array to a double and add to list
            ratings.add(Double.parseDouble(column[2]));
         }
         
         read.close();
         
      } catch (FileNotFoundException e) {
         System.out.println("ERROR - File " + fName + " not found");
      }

      return ratings;
   }
   
    /**
     * Displays a histogram. Reorders a list to sequential order, displays 
     * the range of values in the list from the 0.1 less than the smallest
     * value to 0.1 more than the largest value. For each value, displays
     * how many are found in the list, denoting with asterisks.
     * 
     * @param list
     *            the list to be used for displaying the histogram
     */
   public static void displayHistogram(List<Double> list) {
      
      //place contents of list in sequential order
      Collections.sort(list);
      
      System.out.println("Histogram of Amazon Bestseller Ratings");
      System.out.print("--------------------------------------");
      
      //display each unique value of list once along with smallest value - 0.1 and largest value + 0.1
      for (double val = list.get(0)-0.1; val <= list.get(list.size()-1)+0.1; val+=0.1) {
         //format output to one decimal place
         System.out.printf("\n%.1f ", val);
         
         for (int idx = 0; idx < list.size(); ++idx) {
            //compare displayed double values and values from list, print "*" for every equal value
            if ((list.get(idx) < val + 0.001) && (list.get(idx) > val - 0.001)) {
               System.out.print("*");
            }
         }
      }
      System.out.println("\n--------------------------------------");
   }

    /**
     * Sorts a list, finds the median of its contents.
     * 
     * @param list
     *            the list to be used for finding the median
     * @return a double value representing the computed median
     */
   public static double getMedian(List<Double> list) {
      
      //reorder list in sequential order before calculating
      Collections.sort(list);
      double median;
      
      //different ways to compute median based on whether the number of values is even or odd
      if (list.size()%2 == 0) {
         median = (list.get(list.size() / 2) + list.get((list.size() + 1) / 2)) / 2.0;
      } else {
         //floating-point division is not necessary if the number of terms is odd
         median = list.get((list.size() + 1) / 2);
      }

      return median;
   }

    /**
     * Computes the average of a list by calculating the sum of its
     * contents and then diving it by the number of values it contains.
     * 
     * @param list
     *            the list to be used for finding the average
     * @return a double value representing the computed average
     */
   public static double getAverage(List<Double> list) {
      
      //compute the sum of all contents of list
      double sum = 0.0;
      for (int idx = 0; idx < list.size(); ++idx) {
         sum += list.get(idx);
      }
      
      //divide sum by number of values for average
      double average = sum / list.size();

      return average;
   }
   
    /**
     * Sums the squared difference between each value in a list and the
     * average, divides that number by the total number of values, then
     * takes the square root of it.
     * 
     * @param list
     *            the list to be used for finding the standard deviation
     * @param average
     *            the priorly computed average used for the calculation
     * @return a double value representing the standard deviation
     */
   public static double getStDev(List<Double> list, double average) {
      
      //compute sum of difference between (value at given index and average) squared
      double sumVals = 0.0;
      for (int idx = 0; idx < list.size(); ++idx) {
         //use Math.pow from Java API to square the difference between each value and average
         sumVals += Math.pow(list.get(idx) - average, 2);
      }
      
      //use Math.sqrt to find square root of the sum divided by number of values
      double stDev = Math.sqrt(sumVals / list.size());

      return stDev;
   }

   public static void main(String[] args) {
      Scanner scnr = new Scanner(System.in);

      System.out.print("Enter a filename: ");
      String fileName = scnr.nextLine();

      //call getRatings method to initialize Double List
      List<Double> bookRatings = getRatings(fileName);

      //prevents the rest of the program from compiling if invalid file is entered
      if (bookRatings.size() > 0) {
         //call method to display histogram
         displayHistogram(bookRatings);

         //call corresponding methods to calculate needed statistics
         double bookMedian = getMedian(bookRatings);
         double bookAverage = getAverage(bookRatings);
         double bookStDev = getStDev(bookRatings, bookAverage);

         //display statistics
         System.out.println("Total books rated: " + bookRatings.size());
         System.out.printf("Median score: %.1f\n", bookMedian);
         System.out.printf("Average score: %.1f\n", bookAverage);
         System.out.printf("Standard Deviation: %.2f\n", bookStDev);
      }

      scnr.close();
   }
}
