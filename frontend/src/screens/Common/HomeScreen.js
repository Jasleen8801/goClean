import React from "react";
import { ScrollView, StyleSheet, Text } from "react-native";
import { LinearGradient } from "expo-linear-gradient";

export default function HomeScreen() {
  return (
    <LinearGradient
      colors={["#4c669f", "#3b5998", "#192f6a"]}
    >
      <ScrollView style={styles.container}>
        <Text style={styles.Heading}>Hostel Cleaning</Text>
      </ScrollView>
    </LinearGradient>
  );
}

const styles = StyleSheet.create({
  container: {
    marginTop: 50,
    padding: 10,
    height: "100%",
  },
  Heading: {
    fontSize: 30,
    color: "white",
    textAlign: "center",
  },
});